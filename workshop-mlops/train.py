import torch
import torch.optim as optim
import torch.nn.functional as F
import matplotlib.pyplot as plt

from network import Network
from loadData import train_loader, test_loader

print(len(train_loader))
print(len(next(iter(train_loader))))
print(next(iter(train_loader))[1])
print(next(iter(train_loader))[0].size())
# exit()
device = torch.device('cuda') if torch.cuda.is_available() else torch.device('cpu')

EPOCHS = 5
LEARNING_RATE = 0.01

BATCH_SIZE = 8
TRAINING_SIZE = len(iter(train_loader)) * BATCH_SIZE
TESTING_SIZE = len(iter(test_loader)) * BATCH_SIZE

network = Network()
optimizer = optim.Adam(network.parameters(), lr=LEARNING_RATE)


training_losses = []
training_accuracies = []

testing_losses = []
testing_accuracies = []

def get_num_correct(preds, labels):
    return preds.argmax(dim=1).eq(labels).sum().item()

def train():
    network.train()
    correct_in_episode = 0
    episode_loss = 0

    for batch in train_loader:
        images, labels = batch

        predictions = network(images)
        loss = F.cross_entropy(predictions, labels)

        optimizer.zero_grad()
        loss.backward()
        optimizer.step()

        episode_loss += loss.item()
        correct_in_episode += get_num_correct(predictions, labels)

    training_losses.append(episode_loss / (len(train_loader) * 64))
    training_accuracies.append(correct_in_episode * 100 / (len(train_loader) * 64))

    print(f"Epoch: {epoch + 1} accuracy: {training_accuracies[-1]:.3f} loss: {training_losses[-1]:.3f}", end="\t")

def test():
    network.eval()
    episode_loss = 0
    correct_in_episode = 0

    with torch.no_grad():
        for batch in test_loader:
            images, labels = batch

            predictions = network(images)
            loss = F.cross_entropy(predictions, labels)

            episode_loss = loss.item()
            correct_in_episode += get_num_correct(predictions, labels)

    testing_losses.append(episode_loss)
    testing_accuracies.append(correct_in_episode * 100 / (len(test_loader) * 64))

    print(f'Validation: Accuracy: {testing_accuracies[-1]:.3f} loss: {testing_losses[-1]:.3f}')

for epoch in range(EPOCHS):
    train()
    test()

fig = plt.figure()

plt.plot(list(range(1, len(training_losses)+1)), training_losses, color='blue')
plt.plot(list(range(1, len(testing_losses)+1)), testing_losses, color='red')

plt.legend(['Train Loss', 'Test Loss'], loc='upper right')
plt.xlabel('number of training examples seen')
plt.ylabel('Loss')
plt.show()

fig = plt.figure()

plt.plot(list(range(1, len(training_accuracies)+1)), training_accuracies, color='blue')
plt.plot(list(range(1, len(testing_accuracies)+1)), testing_accuracies, color='red')

plt.legend(['Train Accuracy', 'Test Accuracy'], loc='upper right')
plt.xlabel('number of training examples seen')
plt.ylabel('Accuracy')
plt.savefig('result.png')
