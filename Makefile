all:
    g++ -c matrix.cpp -o matrix.o
    ar rcs libmatrix.a matrix.o
    go build -o main
