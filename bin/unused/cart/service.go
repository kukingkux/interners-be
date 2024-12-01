package cart

import (
	"fmt"

	"github.com/kukingkux/interners-be/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	postIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quanrtity for the post %d", item.PostID)
		}

		postIds[i] = item.PostID
	}

	return postIds, nil
}

func (h *Handler) createOrder(ps []types.Post, items []types.CartItem, userID int) (int, float64, error) {
	postMap := make(map[int]types.Post)
	for _, post := range ps {
		postMap[post.ID] = post
	}

	// check post if in stock
	if err := checkIfCartIsInStock(items, postMap); err != nil {
		return 0, 0, nil
	}
	// calculate total price\
	totalPrice := calculateTotalPrice(items, postMap)
	// reduce quantity of post
	for _, item := range items {
		post := postMap[item.PostID]
		// post.Quantity -= item.Quantity

		h.postStore.UpdatePost(post)
	}
	// create order
	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some shi",
	})
	if err != nil {
		return 0, 0, err
	}
	// create order items
	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID:  orderID,
			PostID:   item.PostID,
			Quantity: item.Quantity,
			// Price:    postMap[item.PostID].Price,
		})
	}

	return orderID, totalPrice, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, posts map[int]types.Post) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	// for _, item := range cartItems {
	// 	post, ok := posts[item.PostID]
	// 	if !ok {
	// 		return fmt.Errorf("post %d is not available in the store, please refresh your cart", item.PostID)
	// 	}

	// 	if post.Quantity < item.Quantity {
	// 		return fmt.Errorf("post %s is not available in the quantity requested", post.Name)
	// 	}
	// }

	return nil
}

func calculateTotalPrice(cartItems []types.CartItem, posts map[int]types.Post) float64 {
	var total float64

	// for _, item := range cartItems {
	// 	post := posts[item.PostID]
	// 	total += post.Price * float64(item.Quantity)
	// }

	return total
}
