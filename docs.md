# Auth routes

1 - register:
/auth/register - Form data - check models/Users newUser
2 - login:
/auth/login - Form data - email/password
3 - refresh:
/auth refresh - token from cookie
4 - verify email:
/auth/verify-email/{id} - user id
5 - forget password:
/auth/forget-password - json -
6 - reset password:
/auth/reset-password/{token} - json - new password - confirm password
7 - change password:
/auth/change-password - json - old/new password

# Restaurant routes

1 - get restaurant by id:
/restaurants/{id}
2 - put restaurant:
/restaurant
3 - delete:
restaurant/{id}
4 - patch:
{id}/image or cover or theme

# Categories routes

1 - get all categories:
/categories?resid=
2 - get single category:
/categories/id
3 - post:
/categories - form - check postCategory model
4 - put/delet:
/categories/{id} - json - restaurant id
5 - patch image:
/categories/{id}/image

# Products

1 - get all:
/products?categoryid=
2 - get single:
/products/id
3 - post:
/products - form - check postProduct model
4 - put/delet:
/products/{id} - json - restaurant id
5 - patch image:
/products/{id}/image

# subscriptions

1 - verify payment
/subscriptions/verify-payment - json - payment id
