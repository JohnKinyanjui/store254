package order_service

import (
	"eccomerce_apis/internal/app/eccommerce/models"

	"github.com/go-pg/pg/v10"
)

func send_order_email(db *pg.DB, order *models.Order) {
	// var userAuth models.UserAuth
	// var user models.User
	// var product []models.OrderProduct
	// var store models.Store

	// type OrderProduct struct {
	// 	Name       string
	// 	Quantity   string
	// 	Price      string
	// 	TotalPrice string
	// }

	// var orderProducts []OrderProduct

	// err := db.Model(&user).Column("full_name").Where("id = ?", order.UserId).Select()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }

	// err = db.Model(&userAuth).Column("email").Where("user_id = ?", order.UserId).Select()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }

	// err = db.Model(&store).Column("name", "currency", "owner_id").Where("id = ?", order.StoreId).Select()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }

	// err = db.Model(&product).Where("order_id = ?", order.Id).Select()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }

	// for _, pr := range order.JProducts {
	// 	var prod models.Product
	// 	var orProd OrderProduct

	// 	err = db.Model(&prod).Column("name").Where("id = ?", pr.ProductId).Select()
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return
	// 	}

	// 	orProd.Name = prod.Name
	// 	orProd.Quantity = fmt.Sprintf("%d", pr.Quantity)
	// 	orProd.Price = fmt.Sprintf("%s %s", store.Currency, formatNumber(pr.Price))
	// 	orProd.TotalPrice = fmt.Sprintf("%s %s", store.Currency, formatNumber(pr.TotalPrice))

	// 	orderProducts = append(orderProducts, orProd)
	// }

	// sendUserEmail := func() {
	// 	var emailMiddleware middlewares.EmailMiddleware
	// 	mb := middlewares.MailBox{
	// 		Address: mail.Address{Name: "", Address: userAuth.Email},
	// 		Subject: "Order Confirmation",
	// 	}
	// 	emailMiddleware = &mb

	// 	template := struct {
	// 		Name      string
	// 		Store     string
	// 		Currency  string
	// 		TotalCost string
	// 		Products  []OrderProduct
	// 		Zone      models.OrderZone
	// 	}{
	// 		Name:      user.FullName,
	// 		Store:     store.Name,
	// 		Currency:  store.Currency,
	// 		TotalCost: fmt.Sprintf("%s %s", store.Currency, formatNumber(order.TotalCost)),
	// 		Products:  orderProducts,
	// 		Zone:      order.Zone,
	// 	}

	// 	err := emailMiddleware.ParseEmailTemplate("templates/email/order_template.html", template)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return
	// 	}

	// 	err = emailMiddleware.SendGoMail()
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return
	// 	}
	//}

	// sendStoreEmail := func() {
	// 	err = db.Model(&userAuth).Column("email").Where("user_id = ?", store.OwnerId).Select()
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return
	// 	}

	// 	var emailMiddleware middlewares.EmailMiddleware
	// 	mb := middlewares.MailBox{
	// 		Address: mail.Address{Name: "", Address: userAuth.Email},
	// 		Subject: "New Order Received",
	// 	}
	// 	emailMiddleware = &mb

	// 	template := struct {
	// 		Name      string
	// 		Store     string
	// 		Currency  string
	// 		TotalCost string
	// 		Products  []OrderProduct
	// 		Zone      models.OrderZone
	// 	}{
	// 		Name:      user.FullName,
	// 		Store:     store.Name,
	// 		Currency:  store.Currency,
	// 		TotalCost: fmt.Sprintf("%s %s", store.Currency, formatNumber(order.TotalCost)),
	// 		Products:  orderProducts,
	// 		Zone:      order.Zone,
	// 	}

	// 	err := emailMiddleware.ParseEmailTemplate("templates/email/order_store_template.html", template)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return
	// 	}

	// 	err = emailMiddleware.SendGoMail()
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return
	// 	}
	// }

	// sendUserEmail()

	// sendStoreEmail()

}

// func (ods *OrderServiceData) send_order_status_email() {
// 	var userAuth models.UserAuth
// 	err := ods.Config.DB.Model(&userAuth).Column("email").Where("user_id = ?", ods.userId).Select()
// 	if err != nil {
// 		log.Println(err.Error())
// 		return
// 	}

// 	var emailMiddleware middlewares.EmailMiddleware
// 	mb := middlewares.MailBox{
// 		Address: mail.Address{Name: "", Address: userAuth.Email},
// 		Subject: "Order Updated",
// 	}
// 	emailMiddleware = &mb
// 	template := struct {
// 		Id     int
// 		Status string
// 	}{
// 		Id:     ods.Params.Order.Id,
// 		Status: ods.Params.Order.OrderStatus,
// 	}

// 	err = emailMiddleware.ParseEmailTemplate("templates/email/order_update_template.html", template)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return
// 	}

// 	err = emailMiddleware.SendGoMail()
// 	if err != nil {
// 		log.Println(err.Error())
// 		return
// 	}
// }
