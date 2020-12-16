package easy104

import "testing"

func Test_minDepth(t *testing.T) {

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "min depth of binary tree",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: 2,
		},
		{
			name: "min depth of binary tree",
			args: args{
				root: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 5,
								Right: &TreeNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := minDepth(tt.args.root)
			if actual != tt.want {
				t.Errorf("expected:%d but actual:%d", tt.want, actual)
			}
		})
	}
}
