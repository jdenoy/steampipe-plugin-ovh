connection "ovh" {
    plugin = "francois2metz/ovh"

    # Go to https://eu.api.ovh.com/createToken/ to create your application key,
    # secret and the consumer key
    # For the rights, GET with the path *
    application_key = ""
    application_secret = ""
    consumer_key = ""

    # OVH Endpoint
    # 'ovh-eu' for OVH Europe API
    # 'ovh-us' for OVH US API
    # 'ovh-ca' for OVH Canada API
    # 'soyoustart-eu' for So you Start Europe API
    # 'soyoustart-ca' for So you Start Canada API
    # 'kimsufi-eu' for Kimsufi Europe API
    # 'kimsufi-ca' for Kimsufi Canada API
    endpoint = "ovh-eu"
}
