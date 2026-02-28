import time

start = time.time() * 1000

i = 0
while i < 1_000_000_000:
    i += 1

end = time.time() * 1000
elapsed = end - start
print(f"Counted to 1 billion in: {elapsed:.0f}ms")
print(f"That is: {elapsed/1000:.1f} seconds")
