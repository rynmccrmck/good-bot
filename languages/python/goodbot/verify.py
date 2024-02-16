import ipaddress
import re
import socket
from aslookup import get_as_data

from .data_loader import load_bot_data

def get_domain_name(ip_address):
    """Get the domain name from the IP address."""
    try:
        hostname = socket.gethostbyaddr(ip_address)[0]
        return hostname
    except socket.herror:
        return "No domain name found"

def get_asn(ip_address):
    """Get the ASN from the IP address."""
    return get_as_data(ip_address).asn

def is_verified_ip(ip, sources, method='dns'):
    """Verifies if the IP matches the domain's resolved IPs."""
    if method == 'dnsReverseForward':
        hostname = get_domain_name(ip)
        if hostname in sources:
            return True
    elif method == 'uaAsnMatch':
        asn = get_asn(ip)
        if asn in sources:
            return True
    elif method == 'uaCidrMatch':
        ipaddr = ipaddress.ip_address(ip)
        for source in sources:
            if ipaddr in ipaddress.ip_network(source):
                return True
    else:
        raise ValueError("Unknown method.")
    return False


def is_user_agent_match(user_agent, ua_pattern):
    """Check if the user agent matches the pattern."""
    return bool(re.match(ua_pattern, user_agent))


def is_good_bot(user_agent, ip_address, bots_data):
    """Check if the user agent and IP address belong to a good bot."""
    for bot in bots_data:
        if is_user_agent_match(user_agent, bot.get('UserAgentPattern')):
            sources = bot['ValidDomains']
            if is_verified_ip(ip_address, sources, bot['Method']):
                return True, bot['name']
    return False, None


if __name__ == '__main__':
    bots_data = load_bot_data()['bots']
    user_agent = "facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)"
    ip_address = "179.60.192.36"
    is_bot, bot_name = is_good_bot(user_agent, ip_address, bots_data)
    if is_bot:
        print(f"Good bot detected: {bot_name}")
    else:
        print("No good bot detected.")
