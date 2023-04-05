REPORT 7:
The Team members: Serik Onbolsyn

Progress:
In this week I started comment part
1. ERD:
 We created a table for comment section like:
   
id : the id of a comment 

   user_id : the id of the user making the comment
   
object_table_name : the table where the commented object is
   
object_id : the id of the commented object in the object_table_name table.
   
text : the text
   
date : the date
   
When we search item or choose them it will bring us to the page of one specific item with the comment attached to the object and will show who wrote this and when 

2. Filtering:
   
We will filter our items by price and quantity , also we will divide our shoes for seasons.
To do it we will just use regular page but we will just order queries when we store our data in array of ProductModel for 2 functions(price,quantity)
