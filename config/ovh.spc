connection "ovh" {
    plugin = "francois2metz/ovh"

    # Go to https://www.ovh.com/auth/api/createToken to create your application key,
    # secret and the consumer key
    # For the rights, GET with the path *
    # application_key = "your-application-key"
    # application_secret = "your-application-secret"
    # consumer_key = "your-consumer-key"

    # Credentials can also be provided via environment variables:
    # - OVH_APPLICATION_KEY
    # - OVH_APPLICATION_SECRET
    # - OVH_CONSUMER_KEY
    # - OVH_ENDPOINT

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
