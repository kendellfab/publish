Page Data
=========
There are just a few things to know about data in publish to get started.  There is a small set of data structures that are used & available to the templates.  For any given template there are a combination of data structures that are sent to that template.

## Data Structures
- Payload
- Post
- Page
- Series
- Time Series

## Pages with Data
There is a small set of pages that will receive data from the backend.  This is a list of pages and the data available to the templates.

### Every Page
- App Config
- All Categories (minus actual posts)
- Recent Posts (10)
- Pages

### Index
- N Posts
- Paginator

### Page
- Page

### Post
- Post
- Series (Where applicable)
- Error

### Series All
- All Series
- Error

### Series One
- Single Series
- Error

### Series Time
- Year
	- Array of Time Series
- Year/Month
	- Array of Time Series (Single series however)
	- Error