pipetail is a utility that does something similar to the following:

    command_with_verbose_output > log.txt &
    while true;do tail -n 20 log.txt;sleep 1;done
    rm log.txt

Example usage:

    find / | ./pipetail -i 2 -n 10
