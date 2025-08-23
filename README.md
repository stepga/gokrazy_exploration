# open questions

- how does /perm work?
  - [x] create/delete/move files dynamically in code (should this be possible?) ... see perhaps mkfs
    - this is possible ... i just fucked up the paths (and `gokrazy/mkfs` might be needed in some cases)
  - [x] add static files into /perm during `gok run`/`gok update`
    - `ExtraFilePaths`/`ExtraFileContents` in a `PackageConfig`
      - cannot write extra files to user-controlled /perm partition
      - other paths within the squashfs (mounted read-only as `/`) are okay

- start services/programs upons start (e.g. `breakglass`)

- disable/guard web gui

- advanced topics
  - build own (even more minimalistic) kernel
  - build image and run in qemu or even emulated on dev machine (faster?)
