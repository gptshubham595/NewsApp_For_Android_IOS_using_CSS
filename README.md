# NewsApp_For_Android_IOS_using_CSS

## Folder 1 contains
 - News App that can show news to mobile devices like android & IOS using HTML, CSS, JQ, JAVASCRIPT
 - Used Ratchet Api to create this feature styling like IOS
 - To run this you can open ````index.html```` file
 - Js folder contains ````titleId.js```` file this contains data to display in json format
 
## Folder 2 contains
 - Api creation using GOLANG for news delivering system.
 - Golang can be installed fro here https://golang.org/doc/install
     - Golang is very useful for writing light-weight microservices. We currently use it for generating APIs that interact with our front-end applications. 
     - If we want to build a small functional microservice quickly, then Golang is a great tool to use. 
     - It's an easy language for developers to learn quickly.
 - Just run ````go run main.go````
 - <b>Welcome home!</b>   this is the first home base page that is return via golang api
 
# Some REQUESTS that can be handled are 

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
### GET: /articles/search/{query}
  It will List every where found
  
## DELETE AN ARTICLE
###  DELETE: /articles/2 

## UPDATE AN ARTICLE
###  PATCH: /articles/2 
  to update article number 2  PATCH: /articles/1 in BODY>raw {"title":"Write new title for the article","description":"Write here updated description for the article"}  


