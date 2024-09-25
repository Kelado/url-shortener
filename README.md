# URL shortener

Turn long, unwieldy URLs into sleek, snackable links with our URL shortener service! It's like giving your URLs a makeover—compact, fast, and easier to share. No more copy-pasting novels, just quick links that get straight to the point. URLs, but make them fun-sized!  

### Setup 
No complicated setup rituals here—just make sure you’ve got Go installed. That’s it! No need to pull your hair out over dependencies. Go's got you covered.


### Build
With Go, building is as easy as a one-liner. Just run ```make build```, and voilà, you've got yourself an executable. It's basically magic, minus the wand.


### Run
Fire up the service with a simple ```make run```. No extra steps, no drama. Just fast, efficient, and up in a snap.

### Test
Testing? Already built-in. Just run ```make test```, sit back, and let Go flex its muscle. Even your tests run smoother because, well... Go.

### API
| **Method** | **Path**   | **Description**                                      |
|------------|------------|------------------------------------------------------|
| POST       | `/shorten` | Give us a long URL, and we’ll shrink it down for you. |
| GET        | `/{code}`  | Pop in the short code, and we’ll take you to the full URL. |
