extern "C" {
    #include <stdio.h>
    #include <stdlib.h>

    void multiply_matrices(int* a, int* b, int* result, int n) {
        printf("C++ function called\n"); // Debug print

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                result[i * n + j] = 0; // Initialize result cell
                for (int k = 0; k < n; k++) {
                    result[i * n + j] += a[i * n + k] * b[k * n + j];
                }
            }
        }
    }
}