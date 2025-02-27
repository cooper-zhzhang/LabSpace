kill -9 `ps -elf | grep "redis" | awk '{print $4}'`
