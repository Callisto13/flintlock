#!/bin/bash

# WARNING: THIS SCRIPT HAS MUTLIPLE PURPOSES.
# TAKE CARE WHEN EDITING.

# This is a scripted version of https://docs.docker.com/storage/storagedriver/device-mapper-driver/#configure-direct-lvm-mode-manually

set -ex

if [[ $(id -u) != 0 ]]; then
  echo "Run this script as root..." >&2
  exit 1
fi

DISK_NAME="${1:-sdb}"
DISK_PATH="/dev/$DISK_NAME"
POOL_NAME="${1:-flintlock}"
PROFILE="/etc/lvm/profile/$POOL_NAME-thinpool.profile"
CROOT=/var/lib/containerd-dev
DIR="${CROOT}/snapshotter/devmapper"

apt update
apt install -y thin-provisioning-tools lvm2

# create physical volume on block device
if [[ $(pvdisplay) != *"$DISK_PATH"* ]]; then
  pvcreate "$DISK_PATH"
fi

# create volume group on same device
if [[ $(vgdisplay) != *"$POOL_NAME"* ]]; then
  vgcreate "$POOL_NAME" "$DISK_PATH"
fi

if [[ $(lvdisplay) != *"$POOL_NAME"* ]]; then
  # create data logical volume
  lvcreate --wipesignatures y -n thinpool "$POOL_NAME" -l 95%VG

  # create metadata logical volume
  lvcreate --wipesignatures y -n thinpoolmeta "$POOL_NAME" -l 1%VG

  # convert logical volumes to thinpool + pool metadata storage
  lvconvert -y \
    --zero n \
    -c 512K \
    --thinpool "$POOL_NAME"/thinpool \
    --poolmetadata "$POOL_NAME"/thinpoolmeta
fi


# write lvm profile with recommended default values
if [[ ! -f "$PROFILE" ]]; then
cat <<'EOF' >> "$PROFILE"
activation {
  thin_pool_autoextend_threshold=80
  thin_pool_autoextend_percent=20
}
EOF
fi

# apply the lvm profile
if [[ $(lvs) != *"$POOL_NAME"* ]]; then
  lvchange --metadataprofile "$POOL_NAME-thinpool" "$POOL_NAME"/thinpool
fi

try=1
max=5
while [ "$try" -le "$max" ]; do
  # check if lvm profile is monitored
  if [[ $(lvs -o+seg_monitor) == *"not monitored"* ]]; then
    # if 'not monitored`, monitor explicitly
    lvchange --monitor y "$POOL_NAME"/thinpool
    ((try=try+1))

    if [[ "$try" -gt "$max" ]]; then
      echo "could not monitor lvm profile"
      exit 1
    fi

    continue
  fi

  break
done

cat << EOF
#
# Add this to your config.toml configuration file and restart containerd daemon
#
[plugins]
  [plugins.devmapper]
    pool_name = "${POOL_NAME}-thinpool"
    root_path = "${DIR}"
    base_image_size = "10GB"
    discard_blocks = true
EOF
