# Dimensionality-reduction-CUR
# dimensional deduction using CUR with best rank k (k &lt; rank(A)) with minimum Frobenius norm.

Please implement a program (in C/C++/Java/Go) to do dimensional deduction using
CUR with best rank k (k < rank(A)) with minimum Frobenius norm. The rows and
columns are indexed starting from 0, 1, 2, …, etc.
Input Example:
1 1 1 0 0
3 3 3 0 0
4 4 4 0 0
5 5 5 0 0
0 0 0 4 4
0 0 0 5 5
0 0 0 2 2
Random rows: 3, 5
Random columns: 2, 3

Output Example:
Scaled matrix C = [
0.9762 0.0
2.9286 0.0
3.9047 0.0
4.8809 0.0
0.0 4.1569
0.0 5.1962
0.0 2.0785
]
Scaled matrix R = [
3.4017 3.4017 3.4017 0.0 0.0
0.0 0.0 0.0 4.1662 4.1662
]
W = [
3.4017 0.0
0.0 4.1662
]
Sigma = [
4.1662 0.0
0.0 3.4017
]
Matrix U = [
0.2940 0.0
0.0 0.2400
]
CUR = C * U * R = [
0.9762 0.9762 0.9762 0.0 0.0
2.9286 2.9286 2.9286 0.0 0.0
3.9047 3.9047 3.9047 0.0 0.0
4.8809 4.8809 4.8809 0.0 0.0
0.0 0.0 0.0 4.1569 4.1569
0.0 0.0 0.0 5.1962 5.1962
0.0 0.0 0.0 2.0875 2.0875
]
Frobenius Norm(A – CUR) = 0.2253

