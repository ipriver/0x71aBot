from django.db import models
from django.contrib.auth.models import User
from django.utils import timezone


class Account(models.Model):
    user = models.OneToOneField(User)
    bot_count = models.IntegerField(default=0)


class Bot(models.Model):
    account = models.ForeignKey(Account)
    channel_name = models.CharField(max_length=50)
    bot_start_time = models. DateTimeField(default=None)

    def bot_start(self):
        self.bot_start_time = timezone.now()
