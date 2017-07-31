from django.db import models


class User(models.Model):
    login = models.CharField(max_length=50)
    email = models.EmailField(max_length=50)
    password = models.CharField(max_length=50)
    reg_date = models.DateField()


class Account(models.Model):
    user = models.ForeignKey(User)
    bot_count = models.IntegerField(default=0)
    bot_list = 0


class Bot(models.Model):
    account = models.ForeignKey(Account)
    channel_name = models.CharField(max_length=50)
    bot_start_time = models.DateField()
