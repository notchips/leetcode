/*
 * @lc app=leetcode id=117 lang=cpp
 *
 * [117] Populating Next Right Pointers in Each Node II
 */
/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() {}

    Node(int _val, Node* _left, Node* _right, Node* _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};
*/

class Solution {
   public:
    Node* connect(Node* root) {
        if (root == nullptr) {
            return nullptr;
        }
        Node* cur = root;             // 记录本层链表的当前结点
        Node* headNode = new (Node);  // 下一层带头节点链表的头节点
        Node* pre = headNode;         // 记录下一层链表的前置结点
        for (; cur != nullptr; cur = cur->next) {
            if (cur->left != nullptr) {
                pre->next = cur->left;
                pre = pre->next;
            }
            if (cur->right != nullptr) {
                pre->next = cur->right;
                pre = pre->next;
            }
        }
        pre->next = nullptr;
        connect(headNode->next);
        return root;
    }
};
