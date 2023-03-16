from django.db import models
from django_enum import EnumField
from rest_framework import serializers

# enum ClotheState(NeufA, NeufS, TBE, BE, Satisfaisant)
#
# enum Size(XS, S, M, L, XL)
#
# type Article {
#   id number
#   name string
#   state ClotheState
#   size Size
#   brand string
#   image_url string
#   description string
#   buy_price number
#   sold boolean (default: false)
#   sold_price number
# }
class State(models.TextChoices):
    NEUF_A = "NeufA", "NeufA"
    NEUF_S = "NeufS", "NeufS"
    TBE = "TBE", "TBE"
    BE = "BE", "BE"
    SATISFAISANT = "Satisfaisant", "Satisfaisant"


class Size(models.TextChoices):
    XS = "XS", "XS"
    S = "S", "S"
    M = "M", "M"
    L = "L", "L"
    XL = "XL", "XL"


class Article(models.Model):
    name = models.CharField(max_length=72)
    state = EnumField(State)
    size = EnumField(Size)
    brand = models.CharField(max_length=24)
    image_url = models.URLField(max_length=200)
    description = models.TextField()
    buy_price = models.PositiveIntegerField()
    sold = models.BooleanField(default=False)
    sold_price = models.PositiveIntegerField(blank=True, null=True)


class ArticleSerializer(serializers.Serializer):
    id = serializers.CharField()
    name = serializers.CharField()
    state = serializers.CharField()
    size = serializers.CharField()
    brand = serializers.CharField()
    image_url = serializers.URLField()
    description = serializers.CharField()
    buy_price = serializers.IntegerField()
    sold = serializers.BooleanField()
    sold_price = serializers.IntegerField()


class CreateArticleValidator(serializers.Serializer):
    name = serializers.CharField(required=True)
    state = serializers.CharField(required=True)
    size = serializers.CharField(required=True)
    brand = serializers.CharField(required=True)
    image_url = serializers.URLField(required=True)
    description = serializers.CharField(required=True)
    buy_price = serializers.IntegerField(required=True)
    sold = serializers.BooleanField(required=False)
    sold_price = serializers.IntegerField(required=False)


class UpdateArticleValidator(serializers.Serializer):
    name = serializers.CharField(required=False)
    state = serializers.CharField(required=False)
    size = serializers.CharField(required=False)
    brand = serializers.CharField(required=False)
    image_url = serializers.URLField(required=False)
    description = serializers.CharField(required=False)
    buy_price = serializers.IntegerField(required=False)
    sold = serializers.BooleanField(required=False)
    sold_price = serializers.IntegerField(required=False)