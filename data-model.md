<h1>Database Tables</h1>
<ul>


<li><h3>Users:</h3></li>

name | id (primary key) | email | password | location 
---  | --- | --- | --- |--- 
data type | int64 | varchar(256) | varchar(50) | varchar(50) 


<li><h3>Shopping Lists:</h3></li>

name | id (primary key) | user_id (foreign key) 
---  | --- | --- 
data type | int64 | int64 


<li><h3>Shopping Lists Contents:</h3></li>

name | id (primary key) | shopping_list_id (foreign key) | product_id (foreign key) | date_content_added 
---  | --- | --- | --- |--- 
data type | int64 | int64 | int64 | timestamp 


<li><h3>Clipped Coupons:</h3></li>

name | id (primary key) | shopping_list_id (foreign key) | coupon_id (foreign key)
---  | --- | --- | ---
data type | int64 | int64 | int64 


<li><h3>Coupons:</h3></li>

name | id (primary key) | product_id (foreign key) | measure_unit_id (foreign key) | flyer_id (foreign key) | measurement_number | discounted_price | coupon_description
---  | --- | --- | --- |--- | --- | --- | --- 
data type | int64 | int64 | int | int64 | int | int | text


<li><h3>Favourite Flyers:</h3></li>

name | id (primary key) | user_id (foreign key) | flyer_id (foreign key) | date_favorited
---  | --- | --- | --- | ---
data type | int64 | int64 | int64 | timestamp


<li><h3>Products:</h3></li>

name | id (primary key) | product_name | product_description
---  | --- | --- | ---
data type | int64 | varchar(256) | text 


<li><h3>Retailers:</h3></li>

name | id (primary key) | category_id (foreign key) | retailer_name 
---  |--- | --- | --- 
data type | int64 | int64 | varchar(50) 


<li><h3>Flyers:</h3></li>

name | id (primary key) | retailer_id (foreign key) | start_date | end_date
---  | --- | --- | ---  |--- 
data type | int64 | int64 | timestamp | timestamp

<li><h3>Measurement Units:</h3></li>

name | id (primary key) | unit_name
---  |--- | --- 
data type | int64 | varchar(50) 


<li><h3>Retailer Categories:</h3></li>

name | id (primary key) | category_name
---  |--- | --- 
data type | int64 | varchar(50) 


<li><h3>Stores:</h3></li>

name | id (primary key) | retailer_id (foreign key) | location 
---  |--- | --- | --- 
data type | int64 | int64 | varchar(50) 


</ul>
