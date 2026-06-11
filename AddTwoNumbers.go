package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resp := &ListNode{}
	current := resp
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		current = current.Next
	}

	return resp.Next
}

func main() {
	lista := addTwoNumbers(&ListNode{1,
		&ListNode{0,
			&ListNode{0,
				&ListNode{0,
					&ListNode{0,
						&ListNode{0,
							&ListNode{0,
								&ListNode{0,
									&ListNode{0,
										&ListNode{0,
											&ListNode{0,
												&ListNode{0,
													&ListNode{0,
														&ListNode{0,
															&ListNode{0,
																&ListNode{0,
																	&ListNode{0,
																		&ListNode{0,
																			&ListNode{0,
																				&ListNode{0,
																					&ListNode{0,
																						&ListNode{0,
																							&ListNode{0,
																								&ListNode{0,
																									&ListNode{0,
																										&ListNode{0,
																											&ListNode{0,
																												&ListNode{0,
																													&ListNode{0,
																														&ListNode{1, nil},
																													},
																												},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}})
	for lista != nil {
		fmt.Println(lista.Val)
		lista = lista.Next
	}
}
