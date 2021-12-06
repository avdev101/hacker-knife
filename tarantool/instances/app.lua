package.setsearchroot('/app')

box.cfg{
    background = false,
    listen = '0.0.0.0:3722',
    log = '| tee',
    log_format = 'plain',
}

local function schema_v1()
    local domain = box.schema.space.create('domain')

    domain:format({
        {name='name', type='string'}
    })
    domain:create_index('primary', {
        type='hash',
        parts={'name'}
    })

    local subdomain = box.schema.space.create('subdomain')

    subdomain:format({
        {name='domain', type='string'},
        {name='name', type='string'},
        {name='cname', type='string'},
        {name='is_new', type='boolean'},
    })
    subdomain:create_index('primary', {
        type='hash',
        parts={'name'}
    })
    subdomain:create_index('domain', {
        type='tree',
        parts={'domain'},
        unique=false
    })

end

box.once('schema_v1', schema_v1)

require('app_queue')