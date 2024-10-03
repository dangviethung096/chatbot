# Quy ước lập trình

## Quy ước đặt tên

Đối với ngôn ngữ go: sẽ được quy ước theo chuẩn camel case và tuân theo chuẩn public sẽ viết hoa chữ đầu, private sẽ viết thường chữ đầu. 
Đối với ngôn ngữ javascript: sẽ được quy ước theo chuẩn camel case.
Đối với html, css: các class, id sẽ được quy ước theo chuẩn kebab case.
Đối với tên file hoặc folder: sẽ được đặt theo chuẩn snake case.

Đối với các trường hợp đặc biệt sẽ có quy ước riêng, không phải lúc nào cũng sẽ áp dụng như trên nhưng sẽ có note hoặc lưu ý.
Còn mặc định không có lưu ý thêm: cách đặt tên sẽ tuân theo quy ước trên.

Diễn giải:
Camel case: 
Cụ thể tên của một hàm, một kiểu (type) sẽ được viết liền các từ được phân cách nhau bằng cách viết hoa chữ cái đầu tiên của từ đó (ngoại từ từ đầu tiên sẽ có một số trường hợp ngoại lệ).
Ví dụ:
Tên của một hàm, biến, kiểu (type) sẽ được viết như sau:
- GetProductByName
- GetProductById
- GetProductByCategory
- productID
- productName
- productCategory

Kebab case: 
Tên của một biến, hàm hoặc kiểu sẽ được viết thường toàn bộ và các từ phân cách nhau bởi dấu gạch dưới "-"
Ví dụ:
- product-id
- product-name
- product-category

Snake case:
Tên của một biến, hàm hoặc kiểu sẽ được viết thường toàn bộ và các từ phân cách nhau bởi dấu gạch dưới "_"
Ví dụ:
- product_id
- product_name
- product_category


### Cách đặt tên hàm trong go
Đối với tên hàm ở trong go sẽ được dùng camel case, đặt tên một cách tường minh và dễ hiểu nhất.
Lưu ý: 
- Cố gắng Viết chức năng của hàm làm đúng việc mà tên của hàm biểu thị.
- Ngoài ra nếu được hãy tuân thủ theo nguyên lý Single responsibility (SOLID)

Ví dụ:
Đối với hàm public
- GetQuotationPage
- InsertCustomer
- ListPagingCustomer

Đối với hàm private
- convertCustomerDAOToDTo
- mappingOrderData

### Cách đặt tên kiểu (type struct) trong go
Đối với kiểu trong go sẽ được dùng camel case.
Tên kiểu sẽ tuân theo cơ chế public, private của go
Đặt tên kiểu theo cách đọc xuôi, sao cho dễ hiểu

Ví dụ:
public
- Customer
- CreateCustomerRequest

private
- customerDAO
- scheduleTask

## Indent convention

## Line break convention