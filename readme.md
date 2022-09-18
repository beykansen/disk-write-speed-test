# Disk Write Speed Test
This little program measures your disk drive write speeds. 

Using the ``dd`` command directly, can give unrealistic results because of cache. So I wrote this little program to run 
```dd``` in ```sync``` mode and measure speed with elapsed time between start and finish. For more information, you can check [this answer](https://askubuntu.com/a/226322/616028).

```bash
Usage: ./disk-write-speed-test [--block-size BLOCK-SIZE] [--count COUNT] [--file-path FILE-PATH]

Options:
  --block-size BLOCK-SIZE, -b BLOCK-SIZE
                         block size as kb [default: 1024]
  --count COUNT, -c COUNT
                         iteration count of block size. If count is 1024 and block size is 1024, data size will be 1024 x 1024kb = 1GB [default: 1024]
  --file-path FILE-PATH, -f FILE-PATH
                         test file to run benchmark on. If you want to benchmark another disk you can specify mount point on that [default: ./disk-write-test-file]
  --help, -h             display this help and exit

```