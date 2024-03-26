-- name: GetProducts :many
select * from products;

-- name: CreateProduct :one
insert into products (id, name, description,created_at)
values (?,?,?,?)
returning *;

-- name: DeleteProduct :exec
delete from products
where id = ?;