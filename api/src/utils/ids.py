import uuid

from nanoid import generate


class Ids:
    def generateNano(self, size: int = 12) -> str:
        return generate(size=size)

    def generateUUID(self) -> str:
        return str(uuid.uuid4())
