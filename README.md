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

- config.txt
  - add boot-otions for raspberry's early-stage boot firmware via `BootloaderExtraLines`
  - see here for key-value-pairs/options: https://www.raspberrypi.com/documentation/computers/config_txt.html
  - config setting `dtparam=audio=on` leads to additional options in `/proc/cmdline`:
    ```diff
    $ git diff -U0 --word-diff --no-index -- cmdline cmdline_with_dtparam_audio_on
    --- a//cmdline
    +++ b//cmdline_with_dtparam_audio_on
    coherent_pool=1M 8250.nr_uarts=1 snd_bcm2835.enable_headphones=0 cgroup_disable=memory numa_policy=interleave nvme.max_host_mem_size_mb=0 {+snd_bcm2835.enable_headphones=1 snd_bcm2835.enable_hdmi=1+} bcm2708_fb.fbwidth=0 bcm2708_fb.fbheight=0 bcm2708_fb.fbswap=1 numa=fake=2 system_heap.max_order=0 smsc95xx.macaddr=DC:A6:32:49:23:C9 vc_mem.mem_base=0x3ec00000 vc_mem.mem_size=0x40000000  console=tty1 root=PARTUUID=60c24cc1-f3f9-427a-8199-4f9f2cab0001/PARTNROFF=2 init=/gokrazy/init rootwait panic=10 oops=panic
    ```
