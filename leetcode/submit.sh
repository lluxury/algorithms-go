echo "sed -e '/package leetcode/d'  $1 > /tmp/$1 && leetcode submit /tmp/$1"

sed -e '/package leetcode/d'  766.toeplitz-matrix.go > /tmp/766.toeplitz-matrix.go && leetcode submit /tmp/766.toeplitz-matrix.go
