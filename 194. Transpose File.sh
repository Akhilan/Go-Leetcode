# https://leetcode.com/problems/transpose-file/description/
# Read from the file file.txt and print its transposed content to stdout.
for ((i = 1; ; i++)); do
    line=$(awk -vn=$i '{print $n}' file.txt)
    [[ $line ]] || break
    echo $line
done
