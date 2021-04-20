go run main.go &
printf "\nPress r : hot restart q : quit\n"
read -n 1 key # -s: do not echo input character. -n 1: read only 1 character (separate with space)
if [[ $key == r ]]
then
    source ./quit.sh
    printf "\nRestarting...\n"
    source ./start.sh
elif [[ $key == q ]]
then
    source ./quit.sh
    printf "\nQuitting...\n"
fi


