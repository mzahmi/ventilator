import config


# r.mset({"alarm_status": "j"})
config.r.mset({"alarm_info": "none"})
# r.mset({"alarm_title": "none"})

# print(r.get("alarm_title"))
print(config.r.get("alarm_info"))
# print(r.get("alarm_text"))
