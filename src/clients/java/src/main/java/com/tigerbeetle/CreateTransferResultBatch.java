//////////////////////////////////////////////////////////
// This file was auto-generated by java_bindings.zig
// Do not manually modify.
//////////////////////////////////////////////////////////

package com.tigerbeetle;

import java.nio.ByteBuffer;


public final class CreateTransferResultBatch extends Batch {


    interface Struct {
        int SIZE = 8;

        int Index = 0;
        int Result = 4;
    }

    /**
     * Creates an empty batch with the desired maximum capacity.
     * <p>
     * Once created, an instance cannot be resized, however it may contain any number of elements
     * between zero and its {@link #getCapacity capacity}.
     *
     * @param capacity the maximum capacity.
     * @throws IllegalArgumentException if capacity is negative.
     */
    public CreateTransferResultBatch(final int capacity) {
        super(capacity, Struct.SIZE);
    }

    CreateTransferResultBatch(final ByteBuffer buffer) {
        super(buffer, Struct.SIZE);
    }

    /**
     * @throws IllegalStateException if not at a {@link #isValidPosition valid position}.
     */
    public int getIndex() {
        final var value = getUInt32(at(Struct.Index));
        return value;
    }

    /**
     * @param index
     * @throws IllegalStateException if not at a {@link #isValidPosition valid position}.
     * @throws IllegalStateException if a {@link #isReadOnly() read-only} batch.
     */
    void setIndex(final int index) {
        putUInt32(at(Struct.Index), index);
    }

    /**
     * @throws IllegalStateException if not at a {@link #isValidPosition valid position}.
     */
    public CreateTransferResult getResult() {
        final var value = getUInt32(at(Struct.Result));
        return CreateTransferResult.fromValue(value);
    }

    /**
     * @param result
     * @throws IllegalStateException if not at a {@link #isValidPosition valid position}.
     * @throws IllegalStateException if a {@link #isReadOnly() read-only} batch.
     */
    void setResult(final CreateTransferResult result) {
        putUInt32(at(Struct.Result), result.value);
    }

}

