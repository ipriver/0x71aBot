from django.conf.urls import url
from . import views

app_name = 'botFactory'
urlpatterns = [
    url(r'^$', views.test, name='index'),
    url(r'^register/$', views.register, name='register'),
]
