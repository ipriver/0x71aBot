from django.db import models


class User(models.Model):
    login = models.CharField(max_length=50)
    email = models.EmailField(max_length=50)
    password = models.CharField(max_length=50)
    reg_date = models.DateField()


class Account(models.Model):
    user = models.ForeignKey(User)
    bot_count = models.IntegerField()
    bot_list = 0


class Bot(models.Model):
    channel_name = models.CharField(max_length=50)
