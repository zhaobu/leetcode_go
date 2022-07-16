

typedef struct Node
{
    int Val;
    Node *Next;
} Node;

Node *reverseList(Node *head)
{
    if (head == nullptr)
    {
        return nullptr;
    }

    Node *help = new (Node);
    help->Next = head;

    Node *pre = help;
    Node *cur = head;

    /*
    1, 2, 3, 4, 5
     */

    while (cur != nullptr)
    {
        Node *next = cur->Next;
        cur.Next = pre;
        pre = cur;
        cur = next;
    }
    return pre;
}