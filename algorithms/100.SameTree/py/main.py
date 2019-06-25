# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    # 递归方式, 时间复杂度O(N)， 空间复杂度O(log(N))
    def isSameTree(self, p: TreeNode, q: TreeNode) -> bool:
        if p is None and q is None:
            return True

        if p is None or q is None:
            return False

        if p.val != q.val:
            return False

        return self.isSameTree(p.left, q.left) and \
            self.isSameTree(p.right, q.right)

    #  迭代方式
    def isSameTreeIter(self, p: TreeNode, q: TreeNode) -> bool:

        def check(p, q):
            if p is None and q is None:
                return True

            if p is None or q is None:
                return False

            if p.val != q.val:
                return False
            return True

        deq = deque([(p, q), ])
        while deq:
            a, b = deq.popleft()
            if not check(a, b):
                return False

            if a:
                deq.append(a.left, b.left)
                deq.append(a.right, b.right)

        return True

    def inorder_morris_traversal(self, root: TreeNode):
        cur = root
        pre = None

        # import pdb
        # pdb.set_trace()
        while cur:
            if cur.left is None:
                print(cur.val)
                cur = cur.right
            else:
                pre = cur.left
                while pre:
                    if pre.right is cur:
                        pre.right = None
                        print(cur.val)
                        cur = cur.right
                        break

                    if pre.right is None:
                        pre.right = cur
                        cur = cur.left
                        break

                    pre = pre.right


if __name__ == '__main__':
    r1 = TreeNode(1)
    r2 = TreeNode(2)
    r3 = TreeNode(3)
    r4 = TreeNode(4)
    r5 = TreeNode(5)
    r6 = TreeNode(6)
    r7 = TreeNode(7)
    r8 = TreeNode(8)
    r9 = TreeNode(9)

    r1.left = r2
    r2.left = r3
    r2.right = r4
    r4.left = r5
    r4.right = r6
    r1.right = r7
    r7.right = r8
    r8.left = r9

    s = Solution()
    s.inorder_morris_traversal(r1)
