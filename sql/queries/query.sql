-- name: GetProducts :many
select * from products;

-- name: CreateProduct :exec
insert into products (id, name, description,created_at,product_type_id)
values (?,?,?,?,?);

-- name: GetProduct :one
select * from products
where id = ?;

-- name: GetProductTypes :many
select * from product_types;

-- name: GetProductType :one
select * from product_types
where id = ?;

-- name: DeleteProductType :exec
delete from product_types
where id = ?;

-- name: UpdateProductType :exec
update product_types
set name = ?, description = ?
where id = ?;

-- name: UpdateProduct :exec
update products
set name = ?, description = ?, product_type_id = ?
where id = ?;

-- name: GetProductsByProductType :many
select * from products
where product_type_id = ?;

-- name: CreateProductType :one
insert into product_types (id, name, description,created_at)
values (?,?,?,?)
returning *;

-- name: DeleteProduct :exec
delete from products
where id = ?;