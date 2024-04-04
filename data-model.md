<h1>Database Tables</h1>
<ul>
<li><h3>Users:</h3></li>

name | id | email | password | location | favorite_flyers_ids | shopping_list_id
---  |--- | --- | --- |--- |--- |---
data type | int64 | varchar(256) | varchar(50) | varchar(50) |  []int64 | int64

<li><h3>Shopping Lists:</h3></li>

name | id (primary key) | user_id (foreign key) |  added_products | clipped_coupons_ids 
---  |--- | --- | --- |--- 
data type | int64 | int64 | text | []int64 

<li><h3>Coupons:</h3></li>

name | id | flyer_id | retailer_id 
---  |--- | --- | ---
data type | int64 | int64 | int64 

<li><h3>Retailers:</h3></li>

name | id | retailer_name | category | stores_ids | active_flyers_ids
---  |--- | --- | --- |--- | --- 
data type | int64 | varchar(256) | varchar(50) | []int64 | []int64

<li><h3>Stores:</h3></li>

name | id | retailer_id | location 
---  |--- | --- | --- 
data type | int64 | int64 | varchar(50) 

<li><h3>Flyers:</h3></li>

name | id | retailer_id | expiry_date | coupons_ids
---  | --- | --- | ---  |--- 
data type | int64 | int64 | timestamp | []int64

</ul>
