from django.conf.urls import url
from . import views

app_name = 'botFactory'
urlpatterns = [
    url(r'^$', views.index, name='index'),
    url(r'^register/$', views.register, name='register'),
    url(r'^register/user_register/$', views.user_register,
        name='user_register'),
    url(r'^sign_in/$', views.sign_in, name='sign_in'),
    url(r'^logout/$', views.user_logout, name='sign_in'),
    url(r'^add_new_bot/$', views.add_new_bot, name='add_new_bot'),
    url(r'^del_bot/$', views.del_bot, name='del_bot'),
]
