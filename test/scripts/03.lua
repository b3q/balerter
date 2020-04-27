-- @interval 5s

local db = require('datasource.mysql.mysql1')
local log = require('log')
local h = require('h')

res, err = db.query('SELECT * FROM users')
if err ~= nil then
    log.error('query error: ' .. err)
    return
end

h.print(res)