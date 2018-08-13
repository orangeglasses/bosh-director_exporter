package main

import (
	"net/http"
	"net/http/httptest"
)

const mockReply = `{
    "value": {
        "agent_id": "fakeid",
        "az": "unknown",
        "configuration_hash": "unused-configuration-hash",
        "deployment": "p-bosh",
        "id": "0",
        "index": 0,
        "job": {
            "name": "bosh",
            "release": "",
            "template": "",
            "templates": [
                {
                    "name": "system-metrics-server",
                    "version": "ddcfb05eb088b4f33c197b11db439254196654b12555a50f76ffbd27331856e1"
                },
                {
                    "name": "nats",
                    "version": "e5e9fe31f5ace35c2331d176f93e5c7d5cc7e0eb"
                },
                {
                    "name": "postgres-9.4",
                    "version": "adee026e9fbc353657fac099f246a48e8e757c2f"
                },
                {
                    "name": "director",
                    "version": "7be8b26fc60099c34ddcdffd9b3dec03557a8305"
                },
                {
                    "name": "health_monitor",
                    "version": "51d772eed8fd559d31c3bf4d3db7f1f5bfa2e214"
                },
                {
                    "name": "uaa",
                    "version": "9589a544a71c24c9a2318d9d6f69644b8bbd9747"
                },
                {
                    "name": "credhub",
                    "version": "a145711dc605d2d85f27f1173e6cc064744144c2"
                },
                {
                    "name": "bbr-credhubdb",
                    "version": "89fd3efdc47e492258e6fa0e3fcec9c1c2bd97ca"
                },
                {
                    "name": "user_add",
                    "version": "6d89712e022acaee63f7e79fea8b48997cdf0dbf"
                },
                {
                    "name": "monit",
                    "version": "f47966bcc7c0ecc76660ae57132b603a9e203eb6"
                },
                {
                    "name": "ca_certs",
                    "version": "48424b102b88233b3d7785fba40dc9a41862398d"
                },
                {
                    "name": "database-backup-restorer",
                    "version": "1634f0b13505993a9582545451af973bfc4fd038"
                },
                {
                    "name": "bpm",
                    "version": "f18421d8c21c425e94bdb3b2df00f3eca2daef29"
                },
                {
                    "name": "vsphere_cpi",
                    "version": "51d2584a7eb00b529cfaa7383d9c80f4b773afe2"
                },
                {
                    "name": "login_banner",
                    "version": "83fdcd1fa5df70c0bda867b3b873f6c4fdf64a90"
                },
                {
                    "name": "blobstore",
                    "version": "c64eb72abc320505ed2d2fbf9cb33701d1f16558"
                },
                {
                    "name": "syslog_forwarder",
                    "version": "419a1c49c9f37ef99e5c3d913e7a7ecd29c1cce5"
                },
                {
                    "name": "jmx-bosh-plugin",
                    "version": "892b4fec41ed814e1c845bbede9c801a5f43ca17"
                }
            ],
            "version": ""
        },
        "job_state": "running",
        "name": "bosh",
        "networks": {
            "cf-test-application": {
                "cloud_properties": {
                    "name": "0249-CF-TEST-APP"
                },
                "default": [
                    "dns",
                    "gateway"
                ],
                "dns": [
                    "10.113.50.135",
                    "10.113.80.135"
                ],
                "gateway": "131.237.214.254",
                "ip": "131.237.214.22",
                "netmask": "255.255.255.0",
                "type": "manual"
            }
        },
        "packages": {
            "backup-and-restore-release-golang": {
                "blobstore_id": "6db265a5-8d56-4cea-7858-4445995d3d10",
                "name": "backup-and-restore-release-golang",
                "sha1": "1aa585d2a805ef57da0ac885e9aeecc5a9a3a161",
                "version": "80f704fa0dd4e47c493e5a7e083dac060d18fd3b"
            },
            "blackbox": {
                "blobstore_id": "4bae91a7-4fa8-4527-4f3b-e4c60c0b987c",
                "name": "blackbox",
                "sha1": "478df16997649bb1410b00a0f8335b19f1599017",
                "version": "35d1337c76dd370a4c1b4d1f207372da6dcf1c84"
            },
            "bosh-gcscli": {
                "blobstore_id": "47ec3b00-21f0-4614-6c44-cc4d877773f0",
                "name": "bosh-gcscli",
                "sha1": "11975a746c300054dd539fb119e9e39a8057b34a",
                "version": "fce60f2d82653ea7e08c768f077c9c4a738d0c39"
            },
            "bpm": {
                "blobstore_id": "f083ce68-6334-44fb-50e2-7d5a25cdd14a",
                "name": "bpm",
                "sha1": "f55d3f5fcd9ef215f7eaca41ceeaf202096dcdb0",
                "version": "3fe49cfa0140be3ebd8da4bdcadfa6b84d847e87"
            },
            "bpm-runc": {
                "blobstore_id": "a85e2c20-2053-41bf-6fb0-2e9b20255a1f",
                "name": "bpm-runc",
                "sha1": "525be4066e9ef1b831f198b9c0246e38f10a0840",
                "version": "c0b41921c5063378870a7c8867c6dc1aa84e7d85"
            },
            "configurator": {
                "blobstore_id": "ef7fe430-a2de-4f35-46c6-646dbfea4868",
                "name": "configurator",
                "sha1": "83dfb62a98d6baa98f57eaafd719bea77456f1d4",
                "version": "0d632a3a9b06f3777bea07d61807ca06ece24dee"
            },
            "credhub": {
                "blobstore_id": "a17d4112-7a2e-4f56-7ff5-d71f07ea33ef",
                "name": "credhub",
                "sha1": "cb1431368696684777074a201e1d1702626f39e7",
                "version": "e3d60a289d5fd414e29ee06e7e5f1a6b3802c792"
            },
            "database-backup-restorer": {
                "blobstore_id": "87fa210c-577e-4094-47d7-bedc76327e4e",
                "name": "database-backup-restorer",
                "sha1": "36db5b64722205e282c9bbca7101ca9e8533d4d7",
                "version": "acfd20580fcce2a1626ccb4cb4b0f6f75de1dac6"
            },
            "database-backup-restorer-boost": {
                "blobstore_id": "989b4910-031c-4e54-4976-86bb53d7b321",
                "name": "database-backup-restorer-boost",
                "sha1": "3895918e754747be25a1121300f36b5ebc1edea6",
                "version": "c3aee061423c7de8e1bbe50db88f82379c54edf3"
            },
            "database-backup-restorer-mariadb": {
                "blobstore_id": "e75205fb-50eb-4b8f-6cb5-5e7c079101c5",
                "name": "database-backup-restorer-mariadb",
                "sha1": "775db01cd46505bb6a252baaaac4408dfcf2af57",
                "version": "1838e76a6125c4e4b61d40a629edd5a94f388154"
            },
            "database-backup-restorer-mysql-5.5": {
                "blobstore_id": "64ae37ef-c723-46a6-6bee-71da7218f746",
                "name": "database-backup-restorer-mysql-5.5",
                "sha1": "6ba2251e956df28f8d3dffeeae71e817aa22cffb",
                "version": "1c3e31db5fb4228b50565f0952a3387752fc4a23"
            },
            "database-backup-restorer-mysql-5.6": {
                "blobstore_id": "8e1ca005-8c50-471b-7519-c467e647db59",
                "name": "database-backup-restorer-mysql-5.6",
                "sha1": "3bb603c473d57bc58cbfe5f89b5cfe7a453df518",
                "version": "fd9ea31a8f101a53382f5415250072e61ab075e0"
            },
            "database-backup-restorer-mysql-5.7": {
                "blobstore_id": "3579ec5e-be4d-4356-6be1-7fc610bb712d",
                "name": "database-backup-restorer-mysql-5.7",
                "sha1": "529454c594a2e5db42f09f766a289b0b22b557ae",
                "version": "d56501baf8c468226230eea0db3ec973e6c0319c"
            },
            "database-backup-restorer-postgres-9.4": {
                "blobstore_id": "62644438-d121-4de2-48a9-231203f82d45",
                "name": "database-backup-restorer-postgres-9.4",
                "sha1": "63e76569b7ddee29ef856c477f6bff96d229dfa6",
                "version": "4f46c5bf4646634cdaf2d9bca922c82a561b8638"
            },
            "database-backup-restorer-postgres-9.6": {
                "blobstore_id": "9ea97658-a18d-49a7-4b87-840849afdcfe",
                "name": "database-backup-restorer-postgres-9.6",
                "sha1": "149abc884b21243d024d5715aae5b71b5fab951d",
                "version": "acaa98101616ef46862d5c9f3bf8c5aa358da19b"
            },
            "davcli": {
                "blobstore_id": "bda8095d-188b-48b5-507a-ff39c0eb64a9",
                "name": "davcli",
                "sha1": "7ca6770f1398653d459a2d2829e8a0c7f8a51ea8",
                "version": "f8a86e0b88dd22cb03dec04e42bdca86b07f79c3"
            },
            "director": {
                "blobstore_id": "ed726cc9-4f3a-4bac-62fd-33f4189abc21",
                "name": "director",
                "sha1": "0286b38f4d3da4455aa62a846a75618cc89c3d97",
                "version": "3a4faca77efdabad4c9129e471c7533c8473e4f4"
            },
            "golang": {
                "blobstore_id": "94702039-8206-4808-713b-3dfbb9129550",
                "name": "golang",
                "sha1": "696a338c453c361f5251d0dff8d88b9bf19f64c7",
                "version": "e3ca1c9440c29ad576d633e9ef6a2f7805a5e8b7"
            },
            "golang-1.9-linux": {
                "blobstore_id": "07dc675a-936a-4452-4cac-19b8fb2958e0",
                "name": "golang-1.9-linux",
                "sha1": "478b499888be94e4023b5c68f772dbd4fc7424be",
                "version": "8d6c67abda8684ce454f0bc74050a213456573ff"
            },
            "gonats": {
                "blobstore_id": "e7d59a32-86d5-475b-7bf2-d47fecd4002d",
                "name": "gonats",
                "sha1": "2e0d8b7d387e600c7886e68ee7dfb80da716c22f",
                "version": "73ec55f11c24dd7c02288cdffa24446023678cc2"
            },
            "health_monitor": {
                "blobstore_id": "004e5088-673b-4ab6-501f-2e2ec22f5202",
                "name": "health_monitor",
                "sha1": "3fbf54fbdf59d42f9e96606906223b9a4973debb",
                "version": "9dbb883b0edeec2ca26fc3d50f909072b22c0e16"
            },
            "iso9660wrap": {
                "blobstore_id": "0a526d8b-72f5-4fb7-4547-fe9727058ea0",
                "name": "iso9660wrap",
                "sha1": "c19ef68d18dcb1e69849614fb69c1f03869fc143",
                "version": "82cd03afdce1985db8c9d7dba5e5200bcc6b5aa8"
            },
            "jmx-bosh-plugin": {
                "blobstore_id": "7b73fe65-8ad0-4009-6527-076dfc55cd94",
                "name": "jmx-bosh-plugin",
                "sha1": "c065f25f1149a8da2f118da8c54b9d98914f89bb",
                "version": "a0b728a06dddc32939294033677746ab30dbda9e"
            },
            "libpq": {
                "blobstore_id": "92043c01-4f90-4397-60bd-faaebd8aee23",
                "name": "libpq",
                "sha1": "d8d3762dfc05419d652fc2e3d70382680695cedd",
                "version": "e2414662250d0498c194c688679661e09ffaa66e"
            },
            "lunaclient": {
                "blobstore_id": "758fcc03-c8fa-4a90-4651-5c12da1da923",
                "name": "lunaclient",
                "sha1": "c34ed36c7da74368e8b5bf1df9dd37506f26730c",
                "version": "b922e045db5246ec742f0c4d1496844942d6167a"
            },
            "mysql": {
                "blobstore_id": "ed7b64dd-aeda-4c92-76b9-b56bbd0cc57b",
                "name": "mysql",
                "sha1": "cccebea3b85376cdc16711d2ccb6bbe76cb8bf2f",
                "version": "898f50dde093c366a644964ccb308a5281c226de"
            },
            "nginx": {
                "blobstore_id": "9bd8b825-aee2-46d3-59a0-3691c953694c",
                "name": "nginx",
                "sha1": "0237ecd3de395a6596e870f20776c2b8deff2a18",
                "version": "5a68865452a3bdcc233867edbbb59c1e18658f6b"
            },
            "openjdk_1.8.0": {
                "blobstore_id": "29628157-2aa3-4dca-6045-35b4042c030c",
                "name": "openjdk_1.8.0",
                "sha1": "2c28bd15a6d2af7f71de64065965430d10b60994",
                "version": "c8846344bf802835ce8b1229de8fa2028d06f603"
            },
            "postgres-9.4": {
                "blobstore_id": "08132596-8384-442b-6b68-d7de02dd4695",
                "name": "postgres-9.4",
                "sha1": "681f502d6883ba1dae76f504f2fd28eb4f517144",
                "version": "52b3a31d7b0282d342aa7a0d62d8b419358c6b6b"
            },
            "ruby-2.4-r4": {
                "blobstore_id": "5e92eb5e-4e4a-401a-4767-3bf99e1682a2",
                "name": "ruby-2.4-r4",
                "sha1": "86ba70ab92a7a078ab2b5016c0d18e946a93899a",
                "version": "0cdc60ed7fdb326e605479e9275346200af30a25"
            },
            "s3cli": {
                "blobstore_id": "936bbbf3-2f4e-4fd2-6aa5-7841ee9cd886",
                "name": "s3cli",
                "sha1": "5ef587ed9b9ed9d9ff48893830f688f90be06bfe",
                "version": "3097f27cb9356172c9ae52de945821c4e338c87a"
            },
            "system-metrics-plugin": {
                "blobstore_id": "35b45765-a4a8-4a47-48ec-62c97f07ba22",
                "name": "system-metrics-plugin",
                "sha1": "sha256:5a367e0e8b24afe3a27b8527fdf3c07a385d48ce5f9d0959224f0a19d09b9b6f",
                "version": "a1922076af198fdb885bb698d5326e17e96e48b61ccd15b1b71679166092edc6"
            },
            "system-metrics-server": {
                "blobstore_id": "af14beae-39e7-47ef-4537-0b5a5acb9759",
                "name": "system-metrics-server",
                "sha1": "sha256:4b0eae2ef4828ab8d4ce0527e7c036467952cf16cc6b4cf35e8250250f794c7a",
                "version": "33a82d2f5c533e0cde31b7af379ec9c6dc513ae5e22abef94f6a97da5ee55618"
            },
            "uaa": {
                "blobstore_id": "ef778168-f1ca-4896-4945-d89b1e143481",
                "name": "uaa",
                "sha1": "83c1cee68fcc4f5c5fdb1f6eb4f18573567efff9",
                "version": "235c8efae3f0267a79094ac39f8c86fd9b7825a2"
            },
            "uaa_utils": {
                "blobstore_id": "479338f3-2370-4466-6018-5692612a7452",
                "name": "uaa_utils",
                "sha1": "0e294a06075e936b687aeb3c50d41e4a6d0730df",
                "version": "90097ea98715a560867052a2ff0916ec3460aabb"
            },
            "verify_multidigest": {
                "blobstore_id": "0d252a70-0c00-433b-7697-d2e312ee12b6",
                "name": "verify_multidigest",
                "sha1": "3720558d6946f283a6157b4107d0c02e9c2def57",
                "version": "8fc5d654cebad7725c34bb08b3f60b912db7094a"
            },
            "vsphere_cpi": {
                "blobstore_id": "82cfbc08-3ff6-420f-5fef-e156aaf8d340",
                "name": "vsphere_cpi",
                "sha1": "577ad812dd4ebc737b97a34fe689ee39592b3871",
                "version": "e1a84e5bd82eb1abfe9088a2d547e2cecf6cf315"
            }
        },
        "persistent_disk": 0,
        "processes": [
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 8372,
                    "percent": 0.1
                },
                "name": "system-metrics-server",
                "state": "running",
                "uptime": {
                    "secs": 1111801
                }
            },
            {
                "cpu": {
                    "total": 0.4
                },
                "mem": {
                    "kb": 149144,
                    "percent": 1.8
                },
                "name": "nats",
                "state": "running",
                "uptime": {
                    "secs": 1111800
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 2925108,
                    "percent": 35.7
                },
                "name": "postgres",
                "state": "running",
                "uptime": {
                    "secs": 1111798
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 454012,
                    "percent": 5.5
                },
                "name": "director",
                "state": "running",
                "uptime": {
                    "secs": 1111797
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 54496,
                    "percent": 0.6
                },
                "name": "worker_1",
                "state": "running",
                "uptime": {
                    "secs": 1111755
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 55884,
                    "percent": 0.6
                },
                "name": "worker_2",
                "state": "running",
                "uptime": {
                    "secs": 1111754
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 58376,
                    "percent": 0.7
                },
                "name": "worker_3",
                "state": "running",
                "uptime": {
                    "secs": 1111753
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 60172,
                    "percent": 0.7
                },
                "name": "worker_4",
                "state": "running",
                "uptime": {
                    "secs": 1111752
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 59828,
                    "percent": 0.7
                },
                "name": "worker_5",
                "state": "running",
                "uptime": {
                    "secs": 1111751
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 65572,
                    "percent": 0.8
                },
                "name": "director_scheduler",
                "state": "running",
                "uptime": {
                    "secs": 1111792
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 61340,
                    "percent": 0.7
                },
                "name": "director_sync_dns",
                "state": "running",
                "uptime": {
                    "secs": 1111791
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 20284,
                    "percent": 0.2
                },
                "name": "director_nginx",
                "state": "running",
                "uptime": {
                    "secs": 1111790
                }
            },
            {
                "cpu": {
                    "total": 0.4
                },
                "mem": {
                    "kb": 51320,
                    "percent": 0.6
                },
                "name": "health_monitor",
                "state": "running",
                "uptime": {
                    "secs": 1111789
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 386676,
                    "percent": 4.7
                },
                "name": "uaa",
                "state": "running",
                "uptime": {
                    "secs": 1111789
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 527744,
                    "percent": 6.4
                },
                "name": "credhub",
                "state": "running",
                "uptime": {
                    "secs": 1111731
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 20812,
                    "percent": 0.2
                },
                "name": "blobstore_nginx",
                "state": "running",
                "uptime": {
                    "secs": 1111758
                }
            },
            {
                "cpu": {
                    "total": 0
                },
                "mem": {
                    "kb": 8360,
                    "percent": 0.1
                },
                "name": "blackbox",
                "state": "running",
                "uptime": {
                    "secs": 1111757
                }
            }
        ],
        "properties": {
            "logging": {
                "max_log_file_size": ""
            }
        },
        "rendered_templates_archive": {
            "blobstore_id": "b3b6bcd6-04ca-4bcc-5693-2ddc447b3f52",
            "sha1": "943ae51d3f2784fb05308ad1adcb5588b4c89b09"
        },
        "resource_pool": {},
        "vitals": {
            "cpu": {
                "sys": "0.2",
                "user": "0.9",
                "wait": "0.1"
            },
            "disk": {
                "ephemeral": {
                    "inode_percent": "1",
                    "percent": "4"
                },
                "persistent": {
                    "inode_percent": "1",
                    "percent": "60"
                },
                "system": {
                    "inode_percent": "31",
                    "percent": "40"
                }
            },
            "load": [
                "0.19",
                "0.10",
                "0.10"
            ],
            "mem": {
                "kb": "2658720",
                "percent": "33"
            },
            "swap": {
                "kb": "385612",
                "percent": "5"
            },
            "uptime": {
                "secs": 1112205
            }
        },
        "vm": {
            "name": "vm-ede30614-d2e3-4a0b-8ac3-3280dd064af5"
        }
    }
}`

func MockServer(username, password, response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authUsername, authPassword, ok := req.BasicAuth()
		if !(authUsername == username && authPassword == password && ok) {
			res.WriteHeader(401)
			return
		}

		res.WriteHeader(200)
		res.Write([]byte(response))
	}))
}
