# SellerApp
Assignment for seller app

# Steps to use
1. Clone the repository
   $ git clone https://github.com/ShreyasG-25/SellerApp.git
   $ cd SellerApp

2. Docker build and Run
   $ docker-compose build
   $ docker-compose up

3. Call API
   Url : http://localhost:3001/products
   Query Parameters : url (amazon product page url we need to scrape from)
   Method : Post

   Response: 
   {
	"code": 201,
	"result": {
		"id": "6352ab8cb9f3c71df2626d9a",
		"url": "https://www.amazon.com/Wireless-ACOZYKITTEN-Noiseless-Portable-Ergonomics/dp/B09KQP94WL/ref=psdc_11036491_t1_B088NDL2G1",
		"created_at": "2022-10-21T14:24:12.468278926Z",
		"updated_at": "2022-10-21T14:24:12.468278998Z",
		"product": {
			"name": "Wireless Mouse, ACOZYKITTEN Noiseless Mouse Ergonomics Optical Silent Mouse with 3 Adjustable DPI Levels, Portable 2.4G USB Cordless Computer Mice for PC, Laptop, Desktop, MacBook, Chromebook - Black",
			"imageURL": "https://m.media-amazon.com/images/I/51y1KKLBErL.__AC_SX300_SY300_QL70_ML2_.jpg",
			"description": "Make sure this fits by entering your model number. \n 【Silent Click and Adjustable DPI】Special soundless design for the right and left buttons, make you concentrate on working or playing games without disturbing others. Exquisite metal scroll wheel and adjustable DPI(1600/1200/1000).   【Plug & Play】Tiny wireless receiver conveniently slots into your computer's USB port, taking up minimal space. USB-receiver stays in your PC USB port or stows conveniently inside the wireless mouse when not in use.   【Portable Size & Ergonomic Design】With compact size, ergonomic design make it easy to store in bag for traveling. Frosted surface, this wieless mouse will fit comfortably in your hands, Provide you with a comfortable office experience.   【2.4GHz Stable Connection】2.4GHz wireless transmission technology provides a powerful and reliable connection up to 33ft, the wireless mouse will turn to sleep mode in 10mins of inactivity, and can be activated by clicking any buttons   【Universal Compatibility】Well compatible with Windows7/8/10/XP, Vista7/8, Linux and Mac OS X 10.4 etc. Fits for desktop, laptop, PC, Macbook and other devices. (Pls Note: this mouse is connected by USB receiver, will NOT compatible with Macbook Pro or other devices which only have Type C ports)",
			"price": "$5.43",
			"totalReviews": 29
		}
	}
}


3. Stop the Container
   Ctrl + C to stop the container.
