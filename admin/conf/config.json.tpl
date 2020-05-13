{
    "server_addr": "127.0.0.1:8083",
    "node": 1,
    "upload_dir": "resource/upload",
    "template_dir": "resource/view",
    "resource_dir": "./resource",
    "db": [
        {
            "driver" :"mysql",
            "alias": "default",
            "username": "homestead",
            "password": "secret",
            "database": "god",
            "table_prefix": "f_",
            "host": "127.0.0.1",
            "port": 3306
        }
    ],
    "redis": [
        {
            "alias": "default",
            "password": "",
            "database": 0,
            "host": "127.0.0.1",
            "port": 6379
        }
    ],
    "session":{
        "driver": "redis",
        "password": "",
        "host": "127.0.0.1",
        "port": 6379,
        "key_pairs": "sess_kp"
    },
    "rpc": [
        {
            "service_name": "admin",
            "addr": "127.0.0.1:5580"
        }
    ]
}