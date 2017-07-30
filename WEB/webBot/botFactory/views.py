from django.shortcuts import render
from django.http import HttpResponse


def test(request):
    return HttpResponse('Неверная форма')


def register(request):
    pass