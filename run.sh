for i in $*;
do
    params=" $params $i"
done
go run github.com/nthnn/AnkoWeb $params