<h1>Database Tables</h1>
<ul>


<li><h3>Users:</h3></li>

name | id (primary key) | email | password | location 
---  | --- | --- | --- |--- 
data type | int64 | string | string | string


<li><h3>Shopping Lists:</h3></li>

name | id (primary key) | user_id (foreign key) 
---  | --- | --- 
data type | int64 | int64 


<li><h3>Shopping Lists Contents:</h3></li>

name | id (primary key) | shopping_list_id (foreign key) | product_id (foreign key) 
---  | --- | --- | --- 
data type | int64 | int64 | int64 


<li><h3>Clipped Coupons:</h3></li>

name | id (primary key) | shopping_list_id (foreign key) | coupon_id (foreign key)
---  | --- | --- | ---
data type | int64 | int64 | int64 


<li><h3>Coupons:</h3></li>

name | id (primary key) | product_id (foreign key) | flyer_id (foreign key) | measurement_unit | measurement_number | original_price | discounted_price | coupon_description 
---  | --- | --- | --- |--- | --- | --- | --- | --- 
data type | int64 | int64 | int64 | enum | float32 | float32 | float32 | string


<li><h3>Favourite Flyers:</h3></li>

name | id (primary key) | user_id (foreign key) | flyer_id (foreign key) 
---  | --- | --- | --- 
data type | int64 | int64 | int64 


<li><h3>Products:</h3></li>

name | id (primary key) | category_id | product_name | product_description
---  | --- | --- | --- | ---
data type | int64 | int | string | string 


<li><h3>Product Categories:</h3></li>

name | id (primary key) | category 
---  | --- | --- 
data type | int | string 


<li><h3>Retailers:</h3></li>

name | id (primary key) | retailer_name 
---  |--- | --- 
data type | int64 | string 


<li><h3>Flyers:</h3></li>

name | id (primary key) | retailer_id (foreign key) | start_date | end_date
---  | --- | --- | ---  | --- 
data type | int64 | int64 | date | date


<li><h3>Stores:</h3></li>

name | id (primary key) | retailer_id (foreign key) | location 
---  |--- | --- | --- 
data type | int64 | int64 | string


</ul>
