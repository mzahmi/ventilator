import config


# r.mset({"alarm_status": "j"})
config.r.mset({"alarm_title": "this title"})
config.r.mset({"alarm_status": "none"})
config.r.mset(
    {"alarm_text": "this text is very long and will take up more than the allocated space"})
# r.mset({"alarm_title": "none"})

# print(r.get("alarm_title"))
print(config.r.get("alarm_title"))
print(config.r.get("alarm_text"))
print(config.r.get("alarm_status"))
# print(r.get("alarm_text"))
