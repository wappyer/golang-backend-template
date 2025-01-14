package view

const NameRequestLog = "request_log"
const SqlRequestLog = `create or replace view request_log as
select l.request_id,
       l.user_id,
       l.client_ip,
       l.method,
       l.route,
       l.path,
       l.code,
       ld.req,
       l.message,
       ld.resp,
       l.req_time,
       l.cost
from log as l
         left join log_detail as ld on l.request_id = ld.request_id`
