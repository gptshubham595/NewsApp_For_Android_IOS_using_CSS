# Welcome home!   

## GET ONE ARTICLES
### GET: /articles/2 
  Gives article number 2  

## GET ALL ARTICLES
### GET: /articles 
  Gives all articles  
  
## ADD AN ARTICLE
### POST: /articles 
  Add this article
  with Body>Raw as {"id":"NUMBER","title":"Give your title","description":"Give your Description",} 

## SEARCH
### GET: /articles/search/<query> 
  It will List every where found
  
## DELETE AN ARTICLE
###  DELETE: /articles/2 

## UPDATE AN ARTICLE
###  PATCH: /articles/2 
  to update article number 2  PATCH: /articles/1 in BODY>raw {"title":"Write new title for the article","description":"Write here updated description for the article"}  

