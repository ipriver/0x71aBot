from django.conf.urls import url
from . import views

app_name = 'botFactory'
urlpatterns = [
    url(r'^$', views.index, name='index'),
    url(r'^register/$', views.register, name='register'),
    url(r'^sign_in/$', views.sign_in, name='sign_in'),
    url(r'^logout/$', views.user_logout, name='sign_in'),
]
