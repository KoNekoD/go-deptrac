<?php

declare (strict_types=1);
namespace Qossmic\Deptrac\Contract\Dependency;

use DEPTRAC_INTERNAL\Symfony\Contracts\EventDispatcher\Event;
/**
 * Event triggered after all the dependencies have been flattened.
 *
 * This occurs when all dependencies caused by class inheritance have been resolved.
 */
final class PostFlattenEvent extends Event
{
}
