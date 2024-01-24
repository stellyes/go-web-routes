# Go Web Routing 

This repository acts as a demonstration of Go's built in web routing and request handlers. The codebase also covers middleware, global app configuration, conditional rendering with HTML templates, and package management.

## Installation

Once you've cloned the repository, run `go run cmd/web/*.go` to run all of the files in the `main` package.

Once your site is up and running, go to [http://localhost:4000/about](http://localhost:4000/about) to see that the webpage has yet to log your IP to the session storage.

Once you've seen the conditionally rendered "We can't find your IP" message, visit the [homepage](http://localhost:4000/), and then go back to the [about](http://localhost:4000/about) page to see that your IP was successfully stored in the global app configuration.

[![Screen-Recording2024-01-23at5-15-41-PM-ezgif-com-video-to-gif-converter.gif](https://i.postimg.cc/VNKz5Ydf/Screen-Recording2024-01-23at5-15-41-PM-ezgif-com-video-to-gif-converter.gif)](https://postimg.cc/HJ8qh1wR)

## Author
Ryan England ([Github](https://github.com/stellyes), [LinkedIn](https://www.linkedin.com/in/ryandengland))