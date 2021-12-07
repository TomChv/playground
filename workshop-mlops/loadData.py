import torch
import torchvision

from torchvision import transforms

train_set = torchvision.datasets.FashionMNIST(
  root="./data/FashionMNIST",
  train=True,
  download=True,
  transform=transforms.Compose([
    transforms.ToTensor()
  ])
)

test_set = torchvision.datasets.FashionMNIST(
  root="./data/FashionMNIST",
  train=False,
  download=True,
  transform=transforms.Compose([
    transforms.ToTensor()
  ])
)

train_loader = torch.utils.data.DataLoader(train_set, batch_size=64)
test_loader = torch.utils.data.DataLoader(test_set, batch_size=64)