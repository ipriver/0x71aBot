# -*- coding: utf-8 -*-
# Generated by Django 1.11.3 on 2017-08-01 11:19
from __future__ import unicode_literals

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('botFactory', '0003_account_bot_id'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='account',
            name='bot_id',
        ),
        migrations.AlterField(
            model_name='account',
            name='user',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to=settings.AUTH_USER_MODEL, unique=True),
        ),
        migrations.AlterField(
            model_name='bot',
            name='bot_start_time',
            field=models.DateField(default=None),
        ),
    ]