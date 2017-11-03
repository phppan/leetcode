package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNode struct {
     Val int
     Next *ListNode
 }

 func initListNodeData (data []int) *ListNode {

     rs := &ListNode{ Val:data[0]}
     tmp := rs
 	for i := 1; i < len(data) ; i++ {
 	    rs2 := &ListNode{Val :data[i]}
 	    rs.Next = rs2
        rs = rs2
    }

    return tmp
 }
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l1Len := listNodeLen(l1)
    l2Len := listNodeLen(l2)

    if l1Len == 0 {
        return l1
    }

    if l2Len == 0 {
        return l2
    }

    if l2Len > l1Len {
        l2, l1 = l1, l2
    }

    l1Arr := make([]int, l1Len)
    l2Arr := make([]int, l2Len)

    for i := 0; i < l1Len; i++ {
        if i < l2Len {
            l1Arr[i] = l1Arr[i] + l2Arr[i]
        }
        if l1Arr[i] > 10 {
            l1Arr[i + 1] = l1Arr[i + 1] + 1
            l1Arr[i] = l1Arr[i] % 10
        }
    }

    return initListNodeData(l1Arr)
}

func addTwoNumbersv1(l1 *ListNode, l2 *ListNode) *ListNode {
    l1Len := listNodeLen(l1)
    l2Len := listNodeLen(l2)

    if l1Len == 0 {
        return l1
    }

    if l2Len == 0 {
        return l2
    }

    if l2Len > l1Len {
        l2, l1 = l1, l2
    }

    rs := &ListNode{}
    tmp := rs
    node := l1
    carry := 0
    for {

        if l2 != nil {
            rs.Val =  node.Val + carry + l2.Val
            node = node.Next
            l2 = l2.Next
        }else{
            rs.Val =  node.Val + carry
            node = node.Next
        }

        if rs.Val >= 10 {
            carry = 1
            rs.Val = rs.Val % 10
        }else{
            carry = 0
        }
        if node == nil {
        	if carry > 0 {
                rs2 := &ListNode{Val:carry}
                rs.Next = rs2
                rs = rs2
            }
            break
        }
        rs2 := &ListNode{}
        rs.Next = rs2
        rs = rs2
    }

    return tmp
}

func listNodeLen(node *ListNode) int {
    l1Len := 0
    for {
        if node == nil {
            break
        }

        l1Len++
        node = node.Next
    }

    return l1Len
}